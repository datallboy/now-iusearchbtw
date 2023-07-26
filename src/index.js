import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';
import 'xterm/css/xterm.css';
import './index.css';

let socket;
var resizedId;

const term = new Terminal();
const fitAddon = new FitAddon();
term.loadAddon(fitAddon);

document.addEventListener('DOMContentLoaded', function () {
  fetch('/new')
    .then((response) => response.json())
    .then((json) => {
      const { containerID } = json.data;
      socket = new WebSocket(
        `ws://localhost:2376/containers/${containerID}/attach/ws?stream=true&stdin=true&stdout=true&stderr=true`
      );
      const attachAddon = new AttachAddon(socket, { bidirectional: true });

      term.loadAddon(attachAddon);

      term.open(document.getElementById('xterm-container'));

      fitAddon.fit();

      socket.onopen = (e) => {
        resizeTty(term);
      };

      socket.onclose = (e) => {
        killContainer(containerID);
      };
    });
});

window.addEventListener('resize', function () {
  clearTimeout(resizedId);
  resizedId = setTimeout(resizeTty, 300);
});

const resizeTty = () => {
  if (term && fitAddon) {
    fitAddon.fit();
    socket.send(`stty columns ${term.cols} rows ${term.rows}`);
    socket.send('\u000A');
    socket.send('\u000C');
  }
};

const killContainer = (containerID) => {
  fetch(`/kill?containerID=${containerID}`, { method: 'DELETE' }).then(() =>
    console.log('Killing container')
  );
};
