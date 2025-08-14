<template>
  <div class="view-root"> {{ title }}

    <InnerView />
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

import InnerView from '../components/authx/SignUpBox.vue'

export default {

  components: { InnerView },

  data() {
    return {
      title: "Example View"
    }
  }

}

</script>
