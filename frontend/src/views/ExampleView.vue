<template>
  <div class="view-root"> {{ title }} </div>
</template>

<style>
.view-root {
  border-width: 1px;
  border-color: gray;
  border-style: solid;
}
</style>

<script>

// import TheWelcome from '../components/TheWelcome.vue'

export default {

  data() {
    return {
      title: "Example View"
    }
  }

}

</script>
