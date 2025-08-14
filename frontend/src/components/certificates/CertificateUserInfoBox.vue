<script lang="js">


export default {

    computed: {
        infoText() {
            return this.stringifyInfo(this.modelValue);
        },
    },

    methods: {
        stringifyInfo(info) {
            let str = ''
            for (let key in info) {
                let value = info[key];
                str = str + key + ' = ' + value + '\n';
            }
            return str;
        },
    },

    props: {
        modelValue: Object,
    },
}

</script>

<style lang="css">
.cert-user-info-box {
    background-color: rgb(207, 207, 207);
    border-style: solid;
    border-color: black;
    border-width: 1px;
    padding: 10px;
    width: 100%;
    color: rgb(71, 71, 71);
    overflow: hidden;
    line-height: 120%;
}
</style>

<template>
    <div class="cert-user-info-box">
        <pre>{{ infoText }}</pre>
    </div>
</template>
