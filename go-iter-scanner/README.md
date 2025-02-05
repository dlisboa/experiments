# iter.Seq scanner in two easy functions

A scanner/tokenizer/lexer in two easy functions.

`runes` returns a sequence of runes; that is fed into `tokens` which takes
each rune and decides what to do. Converting the `rune` sequence into a **pull
iterator** allow `tokens()` to control the read position by calling `next()`.

### Output

```
-- code
if true then
	1
else
	2
end
x = 10
π = 3.1416
-- runes
['i']['f'][' ']['t']['r']['u']['e'][' ']['t']['h']['e']['n']['\n']['\t']['1']['\n']['e']['l']['s']['e']['\n']['\t']['2']['\n']['e']['n']['d']['\t']['\t']['\n']['x'][' ']['='][' ']['1']['0']['\n']['π'][' ']['='][' ']['3']['.']['1']['4']['1']['6']['\n']
-- tokens func
tok: {SYM, "if"}
tok: {SYM, "true"}
tok: {SYM, "then"}
tok: {NUM, "1"}
tok: {SYM, "else"}
tok: {NUM, "2"}
tok: {SYM, "end"}
tok: {SYM, "x"}
tok: {EQ, "="}
tok: {NUM, "10"}
tok: {SYM, "π"}
tok: {EQ, "="}
tok: {NUM, "3.1416"}
```
