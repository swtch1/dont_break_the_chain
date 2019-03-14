# dont_break_the_chain
Simple CLI Application for the
[Don't Break The Chain](https://www.writersstore.com/dont-break-the-chain-jerry-seinfeld/) method,
a simple way to make a little, or a lot, of progress toward your goal every day.

### Build
```bash
git clone https://github.com/swtch1/dont_break_the_chain.git
cd dont_break_the_chain
./build.sh
mv ./dbtc /usr/local/bin

dbtc
```

### Usage
Currently only one chain is supported, the default chain, which is created or used when no arguments are
given.  The usage is simple.  Run the app to mark today as done.  Run the app again every time you complete your
daily goal.  The app keeps track of how long you've kept your chain going for.

### TODO
- Support multiple chains.
- Show dates and length of previous completed chains.
- Show a visual calendar of the current chain.

