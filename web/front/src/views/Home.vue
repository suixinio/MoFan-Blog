<template>
  <v-app app>
    <TopBar></TopBar>
    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="3">
            <Nav></Nav>
          </v-col>
          <v-col>
            <v-sheet max-width="60vw" min-height="80vh" rounded="lg">
              <router-view></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer></Footer>
  </v-app>
</template>

<script>
import Footer from '../components/Footer'
import TopBar from '../components/TopBar'
import Nav from '../components/Nav'

export default {

  components: {
    TopBar,
    Footer,
    Nav
  }
}
</script>
