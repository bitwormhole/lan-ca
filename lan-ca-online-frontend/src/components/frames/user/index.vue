<script lang="js">

import { ElButton } from 'element-plus';
import { ElAvatar } from 'element-plus';

import MyNav from './nav.vue'
import { useSessionStore } from '@/stores/sessions'
import { createNewGetter } from '@/js/getters'


const theSessionStore = useSessionStore()
const theGetter = createNewGetter()


function getStringWithGetter(obj, path, sep, default_value) {
    let str = theGetter.get(obj, path, sep, default_value);
    if (str == null || str == '') {
        str = default_value;
    }
    return str;
}


export default {

    components: { MyNav },

    data() {
        return {}
    },

    computed: {
        avatarImageSrc() {
            let def = '/images/default/avatar.png'
            let info = theSessionStore.info;
            return getStringWithGetter(info, 'avatar', null, def);

        },

        nickname() {
            let def = '匿名';
            let info = theSessionStore.info;
            return getStringWithGetter(info, 'nickname', null, def);
        },

    },

    methods: {},

    mounted() {

    },
}
</script>

<style lang="css"></style>

<template>
    <div class="root">

        <div>
            <h1>
                <ElAvatar :src="avatarImageSrc"></ElAvatar> {{ nickname }}@LanCA:
            </h1>
        </div>

        <MyNav />
        <slot></slot>
    </div>
</template>
