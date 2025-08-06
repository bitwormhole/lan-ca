// base64.js
// 提供 base64 的编码,解码

import { Base64 } from 'js-base64'


const encode = function (src) {
    return Base64.encode(src);
}

const decode = function (src) {
    return Base64.decode(src);
}


export default {
    encode, decode,
} 
