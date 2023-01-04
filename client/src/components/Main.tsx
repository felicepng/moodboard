import { Component, createSignal } from 'solid-js';
import Images from './Images';
import Nav from './Nav';
import Search from './Search';

const Main: Component = () => {
  const [isLoading, setIsLoading] = createSignal(false);

  const onSearch = (input: string) => {
    setIsLoading(true);
    console.log(input);
  };

  return (
    <div class="bg-white w-screen min-h-screen flex flex-col items-center gap-y-14 md:gap-y-20 2xl:gap-y-24 pb-20">
      <Nav />
      <Search onSearch={onSearch} />
      {isLoading() ? <p class="text-xl">Loading...</p> : <Images />}
    </div>
  );
};

export default Main;
