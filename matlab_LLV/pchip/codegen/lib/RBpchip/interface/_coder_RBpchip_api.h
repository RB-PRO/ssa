/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * _coder_RBpchip_api.h
 *
 * Code generation for function '_coder_RBpchip_api'
 *
 */

#ifndef _CODER_RBPCHIP_API_H
#define _CODER_RBPCHIP_API_H

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
extern void RBpchip(real_T coefs[101]);
extern void RBpchip_api(int32_T nlhs, const mxArray *plhs[1]);
extern void RBpchip_atexit(void);
extern void RBpchip_initialize(void);
extern void RBpchip_terminate(void);
extern void RBpchip_xil_shutdown(void);
extern void RBpchip_xil_terminate(void);

#endif

/* End of code generation (_coder_RBpchip_api.h) */
