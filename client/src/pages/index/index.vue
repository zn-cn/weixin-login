<template>
  <div>
    <p>跳转倒计时：{{countDown}}</p>
  </div>
</template>

<script>
import { mapState, mapActions, mapMutations } from 'vuex';
import { get as getLocalStorage } from '@/util/localStorage';

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
    const status = getLocalStorage('status');
    switch (status) {
      case '2':
      case '3':
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
