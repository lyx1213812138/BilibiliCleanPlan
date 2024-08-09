"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.get = void 0;
// TODO: 用store存储Vgroup数据
var axios_1 = require("axios");
function get(url, data) {
    // 用axios发送请求
    var str = JSON.stringify(data);
    axios_1.default.post(url, str).then(function (res) {
        console.log(res.statusText);
    });
}
exports.get = get;
var VgType;
(function (VgType) {
    VgType[VgType["IsSeason"] = 0] = "IsSeason";
    VgType[VgType["IsUp"] = 1] = "IsUp";
})(VgType || (VgType = {}));
var tmp = {
    list: [
        {
            type: VgType.IsSeason,
            UpId: 439465191,
            SeasonId: 1954425
        },
        {
            type: VgType.IsUp,
            UpId: 8014168
        }
    ]
};
get('http://localhost:12121', tmp);
