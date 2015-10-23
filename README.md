# hack-vm-translator
VM Translator for the Hack platform


The Hack VM translator accept a single item as a command line parameter, `AAA`, where AAA is either a filename with a `.vm` extension containing a VM program or the name of a directory containing one or more `.vm` files.

```console
prompt>translator XXX
```

Subsequently, the translator translates all the files that were given as command line parameter to a single assembly language file named `AAA.asm`.

The assembly conforms to the standard VM-on-Hack mapping.

The translator is comprised of a main program and two modules:
* Parser
The role of this module is to parse .vm file.

* Code Writer
The role of this module is to translate each VM command into Hack assembly.

* Main Program
Constructs a `Parser` to parse the VM input (`.vm`) and `CodeWriter` to generate code into the corresponding output file (`.asm`) as previously explained. If the VM input is a directory, a separate `Parser` is used for each file in the directory, however a single `CodeWriter` is used to handle of the output for all the parsers.