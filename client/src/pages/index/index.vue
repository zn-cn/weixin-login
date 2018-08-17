<template>
  <div>
    <p>跳转倒计时：{{countDown}}</p>
  </div>
</template>

<script>
import { mapState, mapActions, mapMutations } from 'vuex';

export default {
  computed: {
    ...mapState('index', ['countDown', 'redirectURI']),
  },
  methods: {
    ...mapMutations('index', ['DECREMENT']),
    ...mapActions('index', ['initRedirectURI']),
    redirect() {
      if (this.redirectURI !== '') {
        location.href = this.redirectURI;
      } else {
        console.log('no redirectURI');
      }
    },
  },
  created() {
    const status = localStorage.getItem('status');
    switch (status) {
      case '2':
        this.$router.push({ name: 'user' });
        break;
      default:
        this.initRedirectURI();
        break;
    }
  },
  mounted() {
    const intervId = setInterval(() => {
      this.DECREMENT(1);
      if (this.countDown <= 0) {
        clearInterval(intervId);
        this.redirect();
      }
    }, 1000);
  },
};
</script>

<style>
</style>
