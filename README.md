# gopjrt

Go Wrappers for [OpenXLA PjRT](https://github.com/openxla/xla/tree/main/xla/pjrt).

This is originally designed to power [GoMLX](github.com/gomlx/gomlx), but it may be used as a standalone, for lower level access to XLA, and other accelerator use cases.

## Examples

**TODO**

## Links to documentation

* [How to use the PJRT C API? #xla/issues/7038](https://github.com/openxla/xla/issues/7038): discussion of folks trying to use PjRT in their projects. The documentation is still lacking as of this writing.
* [PjRT C API README.md](https://github.com/openxla/xla/blob/main/xla/pjrt/c/README.md): a collection of links to other documents.
* [Public Design Document](https://docs.google.com/document/d/1Qdptisz1tUPGn1qFAVgCV2omnfjN01zoQPwKLdlizas/edit).

## Acknowledgements
This project utilizes the following components from the [OpenXLA project](https://openxla.org/):

* OpenXLA PjRT CPU Plugin: This plugin enables execution of XLA computations on the CPU.
* OpenXLA PjRT CUDA Plugin: This plugin enables execution of XLA computations on NVIDIA GPUs.
We gratefully acknowledge the OpenXLA team for their valuable work in developing and maintaining these plugins.

## Licensing:

The [OpenXLA project](https://openxla.org/), including the CPU and CUDA plugins, is [licensed under the Apache 2.0 license](https://github.com/openxla/xla/blob/main/LICENSE).

The CUDA plugin also utilizes the NVIDIA CUDA Toolkit, which is subject to NVIDIA's licensing terms and must be installed by the user.

For more information about OpenXLA, please visit their website at (openxla.org)[https://openxla.org/], or the github page at [github.com/openxla/xla](https://github.com/openxla/xla)
