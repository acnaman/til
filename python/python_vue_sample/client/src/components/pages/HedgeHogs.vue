<template>
  <v-container grid-list-md>
    <v-layout row wrap fill-height>
      <v-flex xs12 sm8 lg4 md5>
        <v-card v-for="item in items" :key="item.id">
          <v-card-title>{{item.name}}</v-card-title>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import router from "../../router/index"
import axios from 'axios';

export default {
  name: 'HedgeHogs',
  data: () => ({
    items: [],
    loading: false,
  }),
  mounted() {
    this.checkLoggedIn();
    this.downloadHedgeHogs();
  },
  methods: {
    checkLoggedIn() {
      this.$session.start();
      if(!this.$session.has("token")) {
        router.push('/auth');
      }
    },
    downloadHedgeHogs() {
      const token = this.$session.get('token')
      console.log(token)
      axios.get('http://localhost:8000/api/hedgehogs/', {
        headers: {
          Authorization: 'JWT ' + token
        }
      }).then((res) => {
        console.log(res.status);
        this.items = res.data;
      })
      
      /*const instance = axios.create({
        baseURL: 'http://localhost:8000/api/hedgehogs/'
      });
      instance.defaults.headers.common['Authorization'] = 'JWT ' + token
      instance.get().then((res) => {
        console.log(res.status);
        this.items = res.data;
      })*/
    }
  }
};
</script>
