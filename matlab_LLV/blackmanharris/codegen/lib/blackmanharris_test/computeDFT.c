/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * computeDFT.c
 *
 * Code generation for function 'computeDFT'
 *
 */

/* Include files */
#include "computeDFT.h"
#include "FFTImplementationCallback.h"
#include "blackmanharris_test.h"
#include "psdfreqvec.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
void computeDFT(const double xin[1024], double varargin_1, creal_T Xx_data[],
                int Xx_size[1], double f_data[], int f_size[1])
{
  double costab_data[513];
  int costab_size[2];
  double sintab_data[513];
  int sintab_size[2];
  int sintabinv_size[2];
  c_FFTImplementationCallback_gen(costab_data, costab_size, sintab_data,
    sintab_size, sintabinv_size);
  Xx_size[0] = 1024;
  c_FFTImplementationCallback_doH(xin, Xx_data, costab_data, sintab_data);
  psdfreqvec(varargin_1, f_data, f_size);
}

/* End of code generation (computeDFT.c) */
