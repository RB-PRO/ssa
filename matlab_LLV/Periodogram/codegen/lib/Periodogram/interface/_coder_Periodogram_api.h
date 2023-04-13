/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * _coder_Periodogram_api.h
 *
 * Code generation for function '_coder_Periodogram_api'
 *
 */

#ifndef _CODER_PERIODOGRAM_API_H
#define _CODER_PERIODOGRAM_API_H

/* Include files */
#include <stddef.h>
#include <stdlib.h>
#include "tmwtypes.h"
#include "mex.h"
#include "emlrt.h"

/* Variable Declarations */
extern emlrtCTX emlrtRootTLSGlobal;
extern emlrtContext emlrtContextGlobal;

/* Function Declarations */
extern void Periodogram(real_T x[1024], real_T output[513]);
extern void Periodogram_api(const mxArray * const prhs[1], int32_T nlhs, const
  mxArray *plhs[1]);
extern void Periodogram_atexit(void);
extern void Periodogram_initialize(void);
extern void Periodogram_terminate(void);
extern void Periodogram_xil_shutdown(void);
extern void Periodogram_xil_terminate(void);

#endif

/* End of code generation (_coder_Periodogram_api.h) */
