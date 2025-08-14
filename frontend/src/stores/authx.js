
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'
import Base64 from '@/js/base64'

const axios_store = useAxiosStore();

export const useAuthxStore = defineStore('authx-store', {

    state() {
        return {
            state1: 1
        }
    },

    getters: {

        getter1(state) {
            return state.state1;
        },

    },

    actions: {

        action1() {
            let url = '/api/v1/example'
            let method = 'GET';
            let self = this;
            axios_store.request({ method, url }).then((res) => {
                // let vo = res.data;
                // let items = vo.domains;
                console.log(res.statusText)
            })
        },

        login({ username, password }) {
            let url = '/api/v1/auth/sign-in'
            let method = 'POST';
            let password_b64 = Base64.encode(password);
            let self = this;
            let auth1 = {
                mechanism: 'password',
                account: username,
                secret: password_b64,
            }
            let data = {
                auth: [auth1]
            }
            let pro = axios_store.request({ method, url, data }).then((res) => {
                // let vo = res.data;
                // let items = vo.domains;
                console.log(res.statusText)
            })
            return pro;
        },

        sign_up() {
            let url = '/api/v1/example'
            let method = 'GET';
            let self = this;
            let pro = axios_store.request({ method, url }).then((res) => {
                // let vo = res.data;
                // let items = vo.domains;
                console.log(res.statusText)
            })
            return pro;
        },

    },

});

// export default { useAuthxStore } 
