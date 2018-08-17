import { GET } from '@/api/ajax';
import { SET_REDIRECT_URI, DECREMENT } from '@/common/mutation-types';
import { GET_REDIRECT_URI } from '@/common/url';

// initial state
const state = {
  countDown: 5,
  redirectURI: '',
};

// getters
const getters = {};
// mutations
const mutations = {
  [SET_REDIRECT_URI](state, payload) {
    state.redirectURI = payload;
  },
  [DECREMENT](state, payload) {
    state.countDown -= payload;
  },
};
// actions
const actions = {
  initRedirectURI({ commit }) {
    GET({
      url: GET_REDIRECT_URI,
      func: (response) => {
        const data = response.data;
        const redirectURI = data.data.redirect_uri;
        commit(SET_REDIRECT_URI, redirectURI);
        localStorage.setItem('status', '2');
      },
      errFunc: (error) => {
        console.log(error);
        const userURL = process.env.NODE_ENV === 'production' ? '/user' : '/#/user';
        commit(SET_REDIRECT_URI, userURL);
        localStorage.setItem('status', '1');
      },
    });
  },
};


export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
