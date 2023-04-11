/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * computepsd.c
 *
 * Code generation for function 'computepsd'
 *
 */

/* Include files */
#include "computepsd.h"
#include "blackmanharris_test.h"
#include "rt_nonfinite.h"
#include "strcmp.h"
#include <string.h>

/* Function Definitions */
void computepsd(const double Sxx1_data[], const double w2_data[], const int
                w2_size[1], const char range[8], double varargout_1_data[], int
                varargout_1_size[1], double varargout_2_data[], int
                varargout_2_size[1], char varargout_3_data[], int
                varargout_3_size[2])
{
  int i;
  int loop_ub;
  static const char cv[10] = { 'r', 'a', 'd', '/', 's', 'a', 'm', 'p', 'l', 'e'
  };

  if (b_strcmp(range)) {
    varargout_1_size[0] = 513;
    varargout_1_data[0] = Sxx1_data[0];
    for (i = 0; i < 511; i++) {
      varargout_1_data[i + 1] = 2.0 * Sxx1_data[i + 1];
    }

    varargout_1_data[512] = Sxx1_data[512];
    varargout_2_size[0] = 513;
    memcpy(&varargout_2_data[0], &w2_data[0], 513U * sizeof(double));
  } else {
    varargout_1_size[0] = 1024;
    memcpy(&varargout_1_data[0], &Sxx1_data[0], 1024U * sizeof(double));
    varargout_2_size[0] = w2_size[0];
    loop_ub = w2_size[0];
    if (0 <= loop_ub - 1) {
      memcpy(&varargout_2_data[0], &w2_data[0], loop_ub * sizeof(double));
    }
  }

  loop_ub = varargout_1_size[0];
  for (i = 0; i < loop_ub; i++) {
    varargout_1_data[i] /= 6.2831853071795862;
  }

  varargout_3_size[0] = 1;
  varargout_3_size[1] = 10;
  for (i = 0; i < 10; i++) {
    varargout_3_data[i] = cv[i];
  }
}

/* End of code generation (computepsd.c) */
