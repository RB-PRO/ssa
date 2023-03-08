/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * pchip.c
 *
 * Code generation for function 'pchip'
 *
 */

/* Include files */
#include "pchip.h"
#include "RBpchip.h"
#include <math.h>

/* Function Definitions */
double exteriorSlope(double d1, double d2, double h1, double h2)
{
  double s;
  double signd1;
  double signs;
  s = ((2.0 * h1 + h2) * d1 - h1 * d2) / (h1 + h2);
  signd1 = d1;
  if (d1 < 0.0) {
    signd1 = -1.0;
  } else if (d1 > 0.0) {
    signd1 = 1.0;
  } else {
    if (d1 == 0.0) {
      signd1 = 0.0;
    }
  }

  signs = s;
  if (s < 0.0) {
    signs = -1.0;
  } else if (s > 0.0) {
    signs = 1.0;
  } else {
    if (s == 0.0) {
      signs = 0.0;
    }
  }

  if (signs != signd1) {
    s = 0.0;
  } else {
    signs = d2;
    if (d2 < 0.0) {
      signs = -1.0;
    } else if (d2 > 0.0) {
      signs = 1.0;
    } else {
      if (d2 == 0.0) {
        signs = 0.0;
      }
    }

    if ((signd1 != signs) && (fabs(s) > fabs(3.0 * d1))) {
      s = 3.0 * d1;
    }
  }

  return s;
}

/* End of code generation (pchip.c) */
