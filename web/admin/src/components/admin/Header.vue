<template>
  <div>
    <a-button type='danger' @click='loginOut'>退出</a-button>
  </div>
</template>

<script>
export default {
  name: 'Header',
  methods: {
    loginOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>

</style>
