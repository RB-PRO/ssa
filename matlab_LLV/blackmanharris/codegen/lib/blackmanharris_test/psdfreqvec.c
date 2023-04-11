/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * psdfreqvec.c
 *
 * Code generation for function 'psdfreqvec'
 *
 */

/* Include files */
#include "psdfreqvec.h"
#include "blackmanharris_test.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
void psdfreqvec(double varargin_4, double w_data[], int w_size[1])
{
  double Fs1;
  double freq_res;
  int i;
  double w1_data[1024];
  if (rtIsNaN(varargin_4)) {
    Fs1 = 6.2831853071795862;
  } else {
    Fs1 = varargin_4;
  }

  freq_res = Fs1 / 1024.0;
  for (i = 0; i < 1024; i++) {
    w1_data[i] = freq_res * (double)i;
  }

  w1_data[512] = Fs1 / 2.0;
  w1_data[1023] = Fs1 - freq_res;
  w_size[0] = 1024;
  memcpy(&w_data[0], &w1_data[0], 1024U * sizeof(double));
}

/* End of code generation (psdfreqvec.c) */
