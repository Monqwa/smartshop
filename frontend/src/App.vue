<template>
  <div id="app">
    <header class="app-header">
      <div class="header-content">
        <div class="logo">
          <router-link to="/">
            <h1 class="header-title">SmartShop</h1>
          </router-link>
        </div>
      </div>
    </header>
    
    <main class="main-content">
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup>

</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  font-family: var(--font-family);
}

.app-header {
  background-color: var(--color-gray-900);
  color: var(--color-white);
  padding: var(--space-md) 0;
  box-shadow: var(--shadow-md);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--space-md);
  display: flex;
  justify-content: center;
  align-items: center;
}

.header-title {
  color: white;
}

.logo h1 {
  font-size: var(--font-size-xl);
  font-weight: 700;
  margin: 0;
  letter-spacing: -0.5px;
}

.main-content {
  flex: 1;
  padding: var(--space-lg) var(--space-md);
  background-color: var(--color-gray-100);
}

</style>
