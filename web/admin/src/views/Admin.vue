<template>
  <a-layout class='container'>
    <Nav></Nav>
    <a-layout>
      <a-layout-header class='headerBtn'>
        <Header></Header>
      </a-layout-header>

      <a-layout-content>
        <router-view :key='$route.fullPath'></router-view>
      </a-layout-content>

      <a-layout-footer class='admin-footer'>
        <Footer></Footer>
      </a-layout-footer>

    </a-layout>
  </a-layout>
</template>

<script>
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'

export default {
  components: {
    Nav: () => import('../components/admin/Nav'),
    Footer,
    Header
  }
}
</script>

<style scoped>
.container {
  height: 100%;
}

.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;

}

.admin-footer {
  padding: 0;
}
</style>
