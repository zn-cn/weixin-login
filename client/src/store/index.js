import Vue from 'vue';
import Vuex from 'vuex';
import user from '@/store/modules/user';
import index from '@/store/modules/index';
import createLogger from '@/util/logger';
import '@/common/common.scss';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    index,
    user,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
});
