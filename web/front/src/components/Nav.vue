<template>
  <v-card class="mx-auto" max-width="300">
    <v-img>
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img :src=profileInfo.avatar alt="">
          </v-avatar>
          <div class="ma-4 grey--text">{{ profileInfo.name }}</div>
        </v-col>
      </v-card-title>
      <v-divider></v-divider>
    </v-img>
    <v-card-title>
      About Me:
    </v-card-title>
    <v-card-text>{{ profileInfo.desc }}</v-card-text>

    <v-divider color="indigo"></v-divider>

    <v-list nav dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ 'mdi-qqchat' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">123123123</v-list-item-content>
      </v-list-item>
    </v-list>

  </v-card>
</template>

<script>
export default {
  name: 'Nav',
  data () {
    return {
      profileInfo: {
        id: 1,
        name: '',
        desc: '',
        img: '',
        avatar: ''
      }
    }
  },
  created () {
    this.getProfileInfo()
  },
  methods: {
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.profileInfo = res.data
      this.profileInfo.id = res.data.ID
    }
  }
}
</script>

<style scoped>

</style>
