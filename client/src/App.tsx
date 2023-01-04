import type { Component } from 'solid-js';
import Images from './components/Images';
import Nav from './components/Nav';
import Search from './components/Search';

const App: Component = () => {
  return (
    <div class="bg-white w-screen min-h-screen flex flex-col items-center gap-y-14 md:gap-y-20 2xl:gap-y-24 pb-20">
      <Nav />
      <Search />
      <Images />
    </div>
  );
};

export default App;
