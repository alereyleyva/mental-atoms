'use client';

import { useState } from 'react';
import { signOut } from 'next-auth/react';

export default function Home() {
  const [isRecording, setIsRecording] = useState<boolean>();
  const [atom, setAtom] = useState<string | null>(null);

  const getRecognition = () => {
    const SpeechRecognition =
      window.SpeechRecognition || window.webkitSpeechRecognition;

    return new SpeechRecognition();
  };

  const handleRecognitionResult = async (event: SpeechRecognitionEvent) => {
    const result: SpeechRecognitionResult = event.results.item(0);

    for (let index = 0; index < result.length; index++) {
      const resultItem = result.item(index);

      setAtom(resultItem.transcript);
    }

    setIsRecording(false);
  };

  const handleRecognition = () => {
    const recognition = getRecognition();

    recognition.onstart = () => setIsRecording(true);
    recognition.onend = () => setIsRecording(false);
    recognition.onresult = handleRecognitionResult;

    recognition.start();
  };

  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      <h1 className={'text-2xl'}>Mental Atoms</h1>
      <button onClick={() => signOut()}>Sign out</button>
      <span>{atom}</span>
      <button onClick={handleRecognition} disabled={isRecording}>
        {isRecording ? 'Stop' : 'Start'} mental atom
      </button>
    </main>
  );
}
