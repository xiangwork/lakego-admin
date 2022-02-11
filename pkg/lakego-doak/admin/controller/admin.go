package controller

import (
    "strings"
    "encoding/json"

    "github.com/deatil/lakego-doak/lakego/tree"
    "github.com/deatil/lakego-doak/lakego/tool"
    "github.com/deatil/lakego-doak/lakego/router"
    "github.com/deatil/lakego-doak/lakego/collection"
    "github.com/deatil/lakego-doak/lakego/support/cast"
    "github.com/deatil/lakego-doak/lakego/support/hash"
    "github.com/deatil/lakego-doak/lakego/support/time"
    "github.com/deatil/lakego-doak/lakego/facade/auth"
    "github.com/deatil/lakego-doak/lakego/facade/config"
    "github.com/deatil/lakego-doak/lakego/facade/cache"
    "github.com/deatil/lakego-doak/lakego/facade/permission"

    "github.com/deatil/lakego-doak/admin/model"
    "github.com/deatil/lakego-doak/admin/model/scope"
    "github.com/deatil/lakego-doak/admin/auth/admin"
    "github.com/deatil/lakego-doak/admin/support/jwt"
    authPassword "github.com/deatil/lakego-doak/lakego/auth/password"
    adminValidate "github.com/deatil/lakego-doak/admin/validate/admin"
    adminRepository "github.com/deatil/lakego-doak/admin/repository/admin"
)

/**
 * 管理员
 *
 * @create 2021-9-2
 * @author deatil
 */
type Admin struct {
    Base
}

/**
 * 列表
 */
func (this *Admin) Index(ctx *router.Context) {
    // 模型
    adminModel := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx))

    // 排序
    order := ctx.DefaultQuery("order", "id__DESC")
    orders := strings.SplitN(order, "__", 2)
    if orders[0] == "" ||
        (orders[0] != "id" &&
        orders[0] != "name" &&
        orders[0] != "last_active" &&
        orders[0] != "add_time") {
        orders[0] = "id"
    }

    if orders[1] == "" || (orders[1] != "DESC" && orders[1] != "ASC") {
        orders[1] = "DESC"
    }

    adminModel = adminModel.Order(orders[0] + " " + orders[1])

    // 搜索条件
    searchword := ctx.DefaultQuery("searchword", "")
    if searchword != "" {
        searchword = "%" + searchword + "%"

        adminModel = adminModel.
            Or("name LIKE ?", searchword).
            Or("nickname LIKE ?", searchword).
            Or("email LIKE ?", searchword)
    }

    // 时间条件
    startTime := ctx.DefaultQuery("start_time", "")
    if startTime != "" {
        adminModel = adminModel.Where("add_time >= ?", this.FormatDate(startTime))
    }

    endTime := ctx.DefaultQuery("end_time", "")
    if endTime != "" {
        adminModel = adminModel.Where("add_time <= ?", this.FormatDate(endTime))
    }

    status := this.SwitchStatus(ctx.DefaultQuery("status", ""))
    if status != -1 {
        adminModel = adminModel.Where("status = ?", status)
    }

    // 分页相关
    start := ctx.DefaultQuery("start", "0")
    limit := ctx.DefaultQuery("limit", "10")

    newStart := cast.ToInt(start)
    newLimit := cast.ToInt(limit)

    adminModel = adminModel.
        Offset(newStart).
        Limit(newLimit)

    list := make([]map[string]interface{}, 0)

    // 列表
    adminModel = adminModel.
        Select([]string{
            "id", "name", "nickname",
            "email", "avatar",
            "is_root", "status",
            "last_active", "last_ip",
            "update_time", "update_ip",
            "add_time", "add_ip",
        }).
        Find(&list)

    var total int64

    // 总数
    err := adminModel.
        Offset(-1).
        Limit(-1).
        Count(&total).
        Error
    if err != nil {
        this.Error(ctx, "获取失败")
        return
    }

    this.SuccessWithData(ctx, "获取成功", router.H{
        "start": start,
        "limit": limit,
        "total": total,
        "list": list,
    })
}

/**
 * 详情
 */
func (this *Admin) Detail(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    var info = model.Admin{}

    // 模型
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Preload("Groups").
        First(&info).
        Error
    if err != nil {
        this.Error(ctx, "账号不存在")
        return
    }

    // 结构体转map
    data, _ := json.Marshal(&info)
    adminData := map[string]interface{}{}
    json.Unmarshal(data, &adminData)

    newInfoGroups:= collection.Collect(adminData["Groups"]).
        Select("id", "parentid", "title", "description").
        ToMapArray()

    avatar := model.AttachmentUrl(adminData["avatar"].(string))

    newInfo := collection.Collect(adminData).
        Only([]string{
            "id", "name", "nickname", "email",
            "is_root", "status",
            "last_active", "last_ip",
            "update_time", "update_ip",
            "add_time", "add_ip",
        }).
        ToMap()

    newInfo["groups"] = newInfoGroups
    newInfo["avatar"] = avatar

    this.SuccessWithData(ctx, "获取成功", newInfo)
}

/**
 * 管理员权限
 */
func (this *Admin) Rules(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    var info = model.Admin{}

    // 模型
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Preload("Groups").
        First(&info).
        Error
    if err != nil {
        this.Error(ctx, "账号不存在")
        return
    }

    // 结构体转map
    data, _ := json.Marshal(&info)
    adminData := map[string]interface{}{}
    json.Unmarshal(data, &adminData)

    groupids := collection.Collect(adminData["Groups"]).
        Pluck("id").
        ToStringArray()

    rules := adminRepository.GetRules(groupids)

    this.SuccessWithData(ctx, "获取成功", router.H{
        "list": rules,
    })
}

/**
 * 添加账号所需分组
 */
func (this *Admin) Groups(ctx *router.Context) {
    adminInfo, _ := ctx.Get("admin")
    adminData := adminInfo.(*admin.Admin)

    list := make([]map[string]interface{}, 0)
    if adminData.IsSuperAdministrator() {
        err := model.NewAuthGroup().
            Order("listorder ASC").
            Order("add_time ASC").
            Select([]string{
                "id",
                "parentid",
                "title",
                "description",
            }).
            Find(&list).
            Error
        if err != nil {
            this.Error(ctx, "获取失败")
            return
        }

        list = collection.
            Collect(list).
            Each(func(item, value interface{}) (interface{}, bool) {
                value2 := value.(map[string]interface{})
                group := map[string]interface{}{
                    "id": value2["id"],
                    "parentid": cast.ToString(value2["parentid"]),
                    "title": value2["title"],
                    "description": value2["description"],
                };

                return group, true
            }).
            ToMapArray()

        newTree := tree.New()
        list2 := newTree.WithData(list).Build("0", "", 1)

        list = newTree.BuildFormatList(list2, "0")
    } else {
        list = adminData.GetGroupChildren()
    }

    this.SuccessWithData(ctx, "获取成功", router.H{
        "list": list,
    })
}

/**
 * 添加
 */
func (this *Admin) Create(ctx *router.Context) {
    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := adminValidate.Create(post)
    if validateErr != "" {
        this.Error(ctx, validateErr)
        return
    }

    status := 0
    if post["status"].(float64) == 1 {
        status = 1
    }

    // 模型
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Where("name = ?", post["name"].(string)).
        Or("email = ?", post["email"].(string)).
        First(&result).
        Error
    if !(err != nil || len(result) < 1) {
        this.Error(ctx, "邮箱或者账号已经存在")
        return
    }

    insertData := model.Admin{
        Name: post["name"].(string),
        Nickname: post["nickname"].(string),
        Email: post["email"].(string),
        Introduce: post["introduce"].(string),
        Status: status,
        AddTime: time.NowTimeToInt(),
        AddIp: tool.GetRequestIp(ctx),
    }

    err2 := model.NewDB().
        Create(&insertData).
        Error
    if err2 != nil {
        this.Error(ctx, "添加账号失败")
        return
    }

    model.NewDB().Create(&model.AuthGroupAccess{
        AdminId: insertData.ID,
        GroupId: post["group_id"].(string),
    })

    this.SuccessWithData(ctx, "添加账号成功", router.H{
        "id": insertData.ID,
    })
}

/**
 * 更新
 */
func (this *Admin) Update(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := adminValidate.Update(post)
    if validateErr != "" {
        this.Error(ctx, validateErr)
        return
    }

    status := 0
    if post["status"].(float64) == 1 {
        status = 1
    }

    // 链接db
    db := model.NewDB()

    // 验证
    result2 := map[string]interface{}{}
    err2 := model.NewAdmin().
        Where(db.Where("id != ?", id).Where("name = ?", post["name"].(string))).
        Or(db.Where("id != ?", id).Where("email = ?", post["email"].(string))).
        First(&result2).
        Error
    if !(err2 != nil || len(result2) < 1) {
        this.Error(ctx, "管理员账号或者邮箱已经存在")
        return
    }

    err3 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "name": post["name"].(string),
            "nickname": post["nickname"].(string),
            "email": post["email"].(string),
            "introduce": post["introduce"].(string),
            "status": status,
        }).
        Error
    if err3 != nil {
        this.Error(ctx, "账号修改失败")
        return
    }

    this.Success(ctx, "账号修改成功")
}

/**
 * 删除
 */
func (this *Admin) Delete(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能删除自己的账号")
        return
    }

    result := map[string]interface{}{}

    // 模型
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    authAdminId := config.New("auth").GetString("Auth.AdminId")
    if authAdminId == id {
        this.Error(ctx, "当前账号不能被删除")
        return
    }

    // 删除
    err2 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Delete(&model.Admin{
            ID: id,
        }).
        Error
    if err2 != nil {
        this.Error(ctx, "账号删除失败")
        return
    }

    this.Success(ctx, "账号删除成功")
}

/**
 * 修改头像
 */
func (this *Admin) UpdateAvatar(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := adminValidate.UpdateAvatar(post)
    if validateErr != "" {
        this.Error(ctx, validateErr)
        return
    }

    err3 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "avatar": post["avatar"].(string),
        }).
        Error
    if err3 != nil {
        this.Error(ctx, "修改头像失败")
        return
    }

    this.Success(ctx, "修改头像成功")
}

/**
 * 修改密码
 */
func (this *Admin) UpdatePasssword(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    password := post["password"].(string)
    if len(password) != 32 {
        this.Error(ctx, "密码格式错误")
        return
    }

    // 生成密码
    pass, encrypt := authPassword.MakePassword(password)

    err3 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "password": pass,
            "password_salt": encrypt,
        }).
        Error
    if err3 != nil {
        this.Error(ctx, "密码修改失败")
        return
    }

    this.Success(ctx, "密码修改成功")
}

/**
 * 启用
 */
func (this *Admin) Enable(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    if result["status"] == 1 {
        this.Error(ctx, "账号已启用")
        return
    }

    err2 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "status": 1,
        }).
        Error
    if err2 != nil {
        this.Error(ctx, "启用账号失败")
        return
    }

    this.Success(ctx, "启用账号成功")
}

/**
 * 禁用
 */
func (this *Admin) Disable(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    if result["status"] == 0 {
        this.Error(ctx, "账号已禁用")
        return
    }

    err2 := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "status": 0,
        }).
        Error
    if err2 != nil {
        this.Error(ctx, "禁用账号失败")
        return
    }

    this.Success(ctx, "禁用账号成功")
}

/**
 * 退出
 */
func (this *Admin) Logout(ctx *router.Context) {
    refreshToken := ctx.Param("refreshToken")
    if refreshToken == "" {
        this.Error(ctx, "refreshToken不能为空")
        return
    }

    c := cache.New()

    if c.Has(hash.MD5(refreshToken)) {
        this.Error(ctx, "refreshToken已失效")
        return
    }

    // jwt
    aud := jwt.GetJwtAud(ctx)
    jwter := auth.NewWithAud(aud)

    // 拿取数据
    claims, claimsErr := jwter.GetRefreshTokenClaims(refreshToken)
    if claimsErr != nil {
        this.Error(ctx, "refreshToken 已失效")
        return
    }

    // 当前账号ID
    refreshAdminid := jwter.GetDataFromTokenClaims(claims, "id")

    // 过期时间
    exp := jwter.GetFromTokenClaims(claims, "exp")
    iat := jwter.GetFromTokenClaims(claims, "iat")
    refreshTokenExpiresIn := exp.(float64) - iat.(float64)

    nowAdminId, _ := ctx.Get("admin_id")
    if refreshAdminid == nowAdminId.(string) {
        this.Error(ctx, "你不能退出你的账号")
        return
    }

    c.Put(hash.MD5(refreshToken), "no", int64(refreshTokenExpiresIn))

    model.NewAdmin().
        Where("id = ?", refreshAdminid).
        Updates(map[string]interface{}{
            "refresh_time": time.NowTimeToInt(),
            "refresh_ip": tool.GetRequestIp(ctx),
        })

    this.Success(ctx, "账号退出成功")
}

/**
 * 授权
 */
func (this *Admin) Access(ctx *router.Context) {
    id := ctx.Param("id")
    if id == "" {
        this.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        this.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx)).
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        this.Error(ctx, "账号信息不存在")
        return
    }

    // 模型
    err2 := model.NewAuthGroupAccess().
        Where("admin_id = ?", id).
        Delete(&model.AuthGroupAccess{}).
        Error
    if err2 != nil {
        this.Error(ctx, "账号授权分组失败")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    access := post["access"].(string)
    if access != "" {
        adminInfo, _ := ctx.Get("admin")
        adminData := adminInfo.(*admin.Admin)

        groupIds := adminData.GetGroupChildrenIds()
        accessIds := strings.Split(access, ",")

        newAccessIds := collection.
            Collect(accessIds).
            Unique().
            ToStringArray()

        intersectAccess := make([]string, 0)
        if !adminData.IsSuperAdministrator() {
            intersectAccess = collection.
                Collect(groupIds).
                Intersect(accessIds).
                ToStringArray()
        } else {
            intersectAccess = newAccessIds
        }

        insertData := make([]model.AuthGroupAccess, 0)
        for _, value := range intersectAccess {
            insertData = append(insertData, model.AuthGroupAccess{
                AdminId: id,
                GroupId: value,
            })
        }

        model.NewDB().Create(&insertData)
    }

    this.Success(ctx, "账号授权分组成功")
}

/**
 * 权限同步
 */
func (this *Admin) ResetPermission(ctx *router.Context) {
    // 清空原始数据
    model.ClearRulesData()

    // 权限
    ruleList := make([]model.AuthRuleAccess, 0)
    err := model.NewAuthRuleAccess().
        Preload("Rule", "status = ?", 1).
        Find(&ruleList).
        Error
    if err != nil {
        this.Error(ctx, "权限同步失败")
        return
    }

    ruleListMap := model.FormatStructToArrayMap(ruleList)

    // 分组
    groupList := make([]model.AuthGroupAccess, 0)
    err2 := model.NewAuthGroupAccess().
        Preload("Group", "status = ?", 1).
        Find(&groupList).
        Error
    if err2 != nil {
        this.Error(ctx, "权限同步失败")
        return
    }

    groupListMap := model.FormatStructToArrayMap(groupList)

    // permission
    cas := permission.New()

    // 添加权限
    if len(ruleListMap) > 0 {
        for _, rv := range ruleListMap {
            rule := rv["Rule"].(map[string]interface{})

            cas.AddPolicy(rv["group_id"].(string), rule["url"].(string), rule["method"].(string))
        }
    }

    // 添加权限
    if len(groupListMap) > 0 {
        for _, gv := range groupListMap {
            cas.AddRoleForUser(gv["admin_id"].(string), gv["group_id"].(string))
        }
    }

    this.Success(ctx, "权限同步成功")
}

