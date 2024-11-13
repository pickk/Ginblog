<template>
  <v-app app>
    <TopBar></TopBar>

    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="12" md="3">
            <ANav></ANav>
          </v-col>
          <v-col cols="12" md="9">
            <v-sheet min-height="80vh" rounded="lg">
              <router-view :key="$route.path"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <AFooter></AFooter>
  </v-app>
</template>

<script>
import TopBar from '../components/TopBar'
import AFooter from '../components/AFooter.vue'
import ANav from '../components/ANav.vue'
export default {

  components: { TopBar, AFooter, ANav }
}
</script>
