<template>
  <div>
    HedgeHog Page
  </div>
</template>

<script>
import router from "../../router/index"
export default {
  name: 'HedgeHogs',
  mounted() {
    this.checkLoggedIn();
  },
  methods: {
    checkLoggedIn() {
      router.push('/auth');
    }
  }
};
</script>
