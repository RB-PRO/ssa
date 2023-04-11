/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * _coder_blackmanharris_test_api.h
 *
 * Code generation for function '_coder_blackmanharris_test_api'
 *
 */

#ifndef _CODER_BLACKMANHARRIS_TEST_API_H
#define _CODER_BLACKMANHARRIS_TEST_API_H

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
extern void blackmanharris_test(real_T spw_j_data[], int32_T spw_j_size[2],
  real_T output_data[], int32_T output_size[2]);
extern void blackmanharris_test_api(const mxArray * const prhs[1], int32_T nlhs,
  const mxArray *plhs[1]);
extern void blackmanharris_test_atexit(void);
extern void blackmanharris_test_initialize(void);
extern void blackmanharris_test_terminate(void);
extern void blackmanharris_test_xil_shutdown(void);
extern void blackmanharris_test_xil_terminate(void);

#endif

/* End of code generation (_coder_blackmanharris_test_api.h) */
