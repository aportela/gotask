<script setup lang="ts">
    import Header from '../components/Header.vue';
</script>

<template>
    <Header />
    <div class="page-wrapper">
        <div class="page-body">
            <div class="container-xl">
                <router-view />
            </div>
        </div>
    </div>
</template>