import { createRouter, createWebHistory } from 'vue-router';
import ListView from './views/ListView.vue';
import CreateListView from './views/CreateListView.vue';

const routes = [
  {
    path: '/',
    name: 'CreateList',
    component: CreateListView
  },
  {
    path: '/list/:listId',
    name: 'ListView',
    component: ListView,
    props: true
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router; 