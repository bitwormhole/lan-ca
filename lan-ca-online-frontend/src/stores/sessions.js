
import { defineStore } from 'pinia'
import { useAxiosStore } from './axios'
import pinia from './pinia'
import { createNewGetter } from '@/js/getters'


const axios_store = useAxiosStore(pinia);

export const useSessionStore = defineStore('session-store', {

    state() {
        return {
            inner_session_info: {
                avatar: '(url)',
                nickname: '(name_string)',
                id: 0,
                email: '(addr)',
            }
        }
    },

    getters: {
        info(state) {
            return state.inner_session_info;
        },
    },

    actions: {

        fetch() {
            let url = '/api/v1/sessions/current'
            let method = 'GET';
            let self = this;
            let getter = createNewGetter();
            axios_store.request({ method, url }).then((res) => {
                let vo = res.data;
                let info = {}
                const name_prefix = 'sessions.0.';
                info.uuid = getter.get(vo, name_prefix + 'uuid');
                info.nickname = getter.get(vo, name_prefix + 'nickname');
                info.avatar = getter.get(vo, name_prefix + 'avatar');
                info.user = getter.get(vo, name_prefix + 'user');
                info.authenticated = getter.get(vo, name_prefix + 'authenticated');
                self.inner_session_info = info;
            })
        },
    },

});

// export default useSessionStore;
