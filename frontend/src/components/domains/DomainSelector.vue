<script lang="js">
// import { useScenesStore } from '@/stores/scenes'
import { useDomainsStore } from '@/stores/domains'
import { ElOption } from 'element-plus';
import { ElSelect } from 'element-plus';
import { ElButton } from 'element-plus';

// const theScenesStore = useScenesStore();
const theDomainStore = useDomainsStore();

export default {

    data() {
        return {
        }
    },

    computed: {

        rawDomainItems() {
            return theDomainStore.items;
        },

        options() {
            let list = this.rawDomainItems;
            return this.wrapItems(list);
        },
    },

    methods: {
        fetchOptions() {
            theDomainStore.fetch()
        },

        wrapItems(list1) {
            let list2 = [];
            for (let i = 0; i < list1.length; i++) {
                list2.push(list1[i]);
            }
            return list2;
        },

        updateThisModelValue(value) {
            this.$emit('update:modelValue', value)
        },
    },


    props: {
        modelValue: Number
    },

}

</script>

<style lang="css">
.domain-selector {
    display: flex;
    flex-direction: row;
    width: 100%;
}

.select-main {
    flex-grow: 1;
}

.button-refresh {
    flex-grow: 0;
    margin-left: 10px;
}
</style>

<template>
    <div class="domain-selector">
        <ElSelect class="select-main" :model-value="modelValue" @update:model-value="updateThisModelValue"
            placeholder="请选择域名">
            <ElOption v-for="item in options" :key="item.id" :label="item.name" :value="item.name"></ElOption>
        </ElSelect>

        <ElButton class="button-refresh" link @click="fetchOptions" title="refresh">
            <ElIcon>
                <Refresh />
            </ElIcon>
        </ElButton>
    </div>
</template>
