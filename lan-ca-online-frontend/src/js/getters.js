// getters.js



function Getter() {
}

Getter.prototype = {

    get: function (object, getting_path, sep, default_value) {
        if (sep == null) {
            sep = '.'
        }
        let str = getting_path + ""
        let path_elements = str.split(sep)
        let p = object;
        for (var i in path_elements) {
            let el_name = path_elements[i];
            if (p == null) {
                break;
            }
            p = p[el_name];
        }
        if (p == null) {
            p = default_value;
        }
        return p;
    }

}



export const createNewGetter = function () {
    return new Getter();
}


// export default createNewGetter

