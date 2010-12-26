set termencoding=utf-8
set encoding=utf-8
set fileencodings=utf-8,cp932,iso-2022-jp,euc-jp
set ambiwidth=double

set nocompatible

set tabstop=4
set expandtab

set list
set listchars=tab:>.

syntax off

set showmatch
set showmode
set hlsearch
set nonumber
set ruler
set laststatus=2
set statusline=%y%{GetStatusEx()}%F%m%r%=<%l:%c>
hi StatusLine term=NONE cterm=NONE ctermfg=black ctermbg=white

set wildmenu
set wildmode=list:longest

set autoread

set nobackup

set magic
set incsearch
set ignorecase
set smartcase
set wrapscan

set history=50

function! GetStatusEx()
let str = ''
let str = str . '' . &fileformat . ']'
if has('multi_byte') && &fileencoding != ''
let str = '[' . &fileencoding . ':' . str
else
let str = '[' . &encoding . ':' . str
endif
return str
endfunction
