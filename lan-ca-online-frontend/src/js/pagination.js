// pagination.js

import { createNewGetter } from './getters'

const theGetter = createNewGetter()


/******
 *  标准化 pagination 对象
 * ***/
function normalizePagination(pagination) {
    const sep = '.'
    let dst = {}
    dst.page = theGetter.get(pagination, 'page', sep, 1);
    dst.size = theGetter.get(pagination, 'size', sep, 5);
    dst.total = theGetter.get(pagination, 'total', sep, 0);
    return dst;
}


export const Pagination = {
    normalizePagination
}
