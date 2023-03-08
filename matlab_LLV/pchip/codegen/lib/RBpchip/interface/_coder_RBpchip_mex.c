/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * _coder_RBpchip_mex.c
 *
 * Code generation for function '_coder_RBpchip_mex'
 *
 */

/* Include files */
#include "_coder_RBpchip_mex.h"
#include "_coder_RBpchip_api.h"

/* Function Declarations */
MEXFUNCTION_LINKAGE void RBpchip_mexFunction(int32_T nlhs, mxArray *plhs[1],
  int32_T nrhs);

/* Function Definitions */
void RBpchip_mexFunction(int32_T nlhs, mxArray *plhs[1], int32_T nrhs)
{
  const mxArray *outputs[1];
  emlrtStack st = { NULL,              /* site */
    NULL,                              /* tls */
    NULL                               /* prev */
  };

  st.tls = emlrtRootTLSGlobal;

  /* Check for proper number of arguments. */
  if (nrhs != 0) {
    emlrtErrMsgIdAndTxt(&st, "EMLRT:runTime:WrongNumberOfInputs", 5, 12, 0, 4, 7,
                        "RBpchip");
  }

  if (nlhs > 1) {
    emlrtErrMsgIdAndTxt(&st, "EMLRT:runTime:TooManyOutputArguments", 3, 4, 7,
                        "RBpchip");
  }

  /* Call the function. */
  RBpchip_api(nlhs, outputs);

  /* Copy over outputs to the caller. */
  emlrtReturnArrays(1, plhs, outputs);
}

void mexFunction(int32_T nlhs, mxArray *plhs[], int32_T nrhs, const mxArray
                 *prhs[])
{
  (void)prhs;
  mexAtExit(&RBpchip_atexit);

  /* Module initialization. */
  RBpchip_initialize();

  /* Dispatch the entry-point. */
  RBpchip_mexFunction(nlhs, plhs, nrhs);

  /* Module termination. */
  RBpchip_terminate();
}

emlrtCTX mexFunctionCreateRootTLS(void)
{
  emlrtCreateRootTLS(&emlrtRootTLSGlobal, &emlrtContextGlobal, NULL, 1);
  return emlrtRootTLSGlobal;
}

/* End of code generation (_coder_RBpchip_mex.c) */
