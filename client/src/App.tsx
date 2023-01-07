import type { Component } from 'solid-js';
import { createResource, createSignal } from 'solid-js';
import { generateImages } from './controllers/image';
import Images from './components/Images';
import Nav from './components/Nav';
import Prompts from './components/Prompts';
import Search from './components/Search';
import Text from './components/Text';

const [theme, setTheme] = createSignal<string[]>(['']);
const [images] = createResource(theme, generateImages);

const App: Component = () => {
  return (
    <div class="bg-white w-screen min-h-screen flex flex-col items-center gap-y-14 md:gap-y-20 2xl:gap-y-24 pb-20">
      <Nav />
      <Search />
      {!images.loading && images().urls.length ? (
        <>
          <Images />
          <Prompts />
        </>
      ) : (
        <Text />
      )}
    </div>
  );
};

export { theme, setTheme, images };
export default App;
