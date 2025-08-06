<template>
  <div class="view-root">
    <InnerView> </InnerView>
  </div>
</template>

<style>
.view-root {
  border-width: 1px;
  border-color: gray;
  border-style: solid;
}
</style>

<script>

import InnerView from '../components/authx/LoginBox.vue'

export default {

  components: { InnerView },

  data() {
    return {
      title: "Login"
    }
  }

}

</script>
