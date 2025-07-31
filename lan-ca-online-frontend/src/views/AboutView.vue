<script>

import { ElTable } from 'element-plus';
import { ElButton } from 'element-plus';


export default {


  data() {
    return {
      items: [
        { id: 1, name: "ak", domain: "ak17" },
        { id: 2, name: "xsf", domain: "xushifu" },
        { id: 3, name: "mlgm", domain: "milegema" }
      ]
    }
  }

}



</script>

<template>
  <div class="about">
    <h1>This is an about page</h1>

    <ElButton> hello, vue3 </ElButton>

    <ElTable :data="items">
      <el-table-column prop="id" label="ID" width="200" />
      <el-table-column prop="name" label="Name" width="200" />
      <el-table-column prop="domain" label="Domain" width="200" />
    </ElTable>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
