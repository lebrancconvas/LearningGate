import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';

export const state = () => ({
  users: [],
  isLogin: false,
  playgroundMessage: "Ayo"
});

export const mutations = {
  ADD_USER(state, user) {
    const newUser = { id: uuidv4(), ...user };
    state.users.push(newUser);
    const url = "http://localhost:8000/api/v1/users";
    axios.post(url, newUser)
      .then(res => {
        console.log(res.data);
      })
      .catch(err => {
        console.error(err);
        console.log(`[ERROR] Posting Data error.`);
      })
  },
  CHANGE_MESSAGE(state, newMessage) {
    state.playgroundMessage = newMessage;
  }
};
