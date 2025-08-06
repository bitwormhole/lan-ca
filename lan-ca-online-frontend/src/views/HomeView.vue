<template>
    <UserFrame>
        <div class="banner-box">
            <HomeBanner />
        </div>
    </UserFrame>
</template>

<style lang="css">
.banner-box {
    color: rgb(3, 56, 3);
    padding: 30px;
    /*
    background-color: rgba(2, 44, 2, 0.637);
    
    border-color: rgb(114, 84, 49);
    border-width: 10px;
    border-style: solid; */
}
</style>

<script setup>
import HomeBanner from '@/components/home/HomeBanner.vue'
</script>
