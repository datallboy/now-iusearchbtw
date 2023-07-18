import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';
import 'xterm/css/xterm.css';

document.addEventListener('DOMContentLoaded', function () {
  fetch('/new')
    .then((response) => response.json())
    .then((json) => {
      const term = new Terminal();
      const { containerID } = json.data;
      console.log(`Connecting to container: ${containerID}`);
      const socket = new WebSocket(
        `ws://localhost:2376/containers/${containerID}/attach/ws?stream=true&stdin=true&stdout=true&stderr=true`
      );
      const attachAddon = new AttachAddon(socket, { bidirectional: true });

      const fitAddon = new FitAddon();

      term.loadAddon(attachAddon);
      term.loadAddon(fitAddon);

      term.open(document.getElementById('xterm-container'));

      fitAddon.fit();

      socket.onopen = (e) => {
        console.log('Connection established');
        socket.send('\u000A');
      };

      socket.onclose = (e) => {
        killContainer(containerID);
      };
    });
});

const killContainer = (containerID) => {
  fetch(`/kill?containerID=${containerID}`, { method: 'DELETE' }).then(() =>
    console.log('Killing container')
  );
};
