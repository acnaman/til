var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello, Vue.js!',
    datalist: ['りんご', 'ばなな', 'いちご'],
    count: 0
  },
  methods: {
    handleClick: function(event) {
      app.datalist.push('newitem')
    }
  }
})