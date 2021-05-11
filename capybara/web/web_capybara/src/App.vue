<template>
<div>
  <router-view></router-view>
  <div id="app">APP 根组件888</div>
  <router-link to="/login">登录页面</router-link>
  <router-link to="/collect">收藏页面</router-link>
</div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
</style>
