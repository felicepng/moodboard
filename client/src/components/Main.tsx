import { Component, createResource, createSignal } from 'solid-js';
import { generateImages } from '../controllers/image';
import Images from './Images';
import Nav from './Nav';
import Prompts from './Prompts';
import Search from './Search';

const Main: Component = () => {
  const [search, setSearch] = createSignal<string>('');
  const [images] = createResource(search, generateImages);

  const onSearch = (input: string) => {
    setSearch(input);
  };

  const displayText = () => {
    return search() ? (
      <p class="text-xl md:text-2xl text-slate-500">loading...</p>
    ) : (
      <p class="text-center text-xl md:text-2xl text-slate-500 px-14 md:px-20">
        press 'search' to generate a moodboard. if a theme is not provided, a
        random moodboard will be generated.
      </p>
    );
  };

  return (
    <div class="bg-white w-screen min-h-screen flex flex-col items-center gap-y-14 md:gap-y-20 2xl:gap-y-24 pb-20">
      <Nav />
      <Search onSearch={onSearch} />
      {!images.loading && images().urls.length > 0 ? (
        <Images urls={images().urls} />
      ) : (
        displayText()
      )}
      {!images.loading && images().prompts.length > 0 && (
        <Prompts prompts={images().prompts} />
      )}
    </div>
  );
};

export default Main;
