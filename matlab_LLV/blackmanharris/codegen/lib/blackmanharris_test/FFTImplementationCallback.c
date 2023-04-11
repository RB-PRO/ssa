/*
 * Academic License - for use in teaching, academic research, and meeting
 * course requirements at degree granting institutions only.  Not for
 * government, commercial, or other organizational use.
 *
 * FFTImplementationCallback.c
 *
 * Code generation for function 'FFTImplementationCallback'
 *
 */

/* Include files */
#include "FFTImplementationCallback.h"
#include "blackmanharris_test.h"
#include "rt_nonfinite.h"
#include <math.h>
#include <string.h>

/* Function Declarations */
static void c_FFTImplementationCallback_get(creal_T y_data[], const creal_T
  reconVar1_data[], const creal_T reconVar2_data[], const int wrapIndex_data[]);

/* Function Definitions */
static void c_FFTImplementationCallback_get(creal_T y_data[], const creal_T
  reconVar1_data[], const creal_T reconVar2_data[], const int wrapIndex_data[])
{
  double temp1_re;
  double temp1_im;
  double y_im;
  double y_re;
  double b_y_im;
  int i;
  int temp2_re_tmp_tmp;
  int temp2_re_tmp;
  double temp2_re;
  double temp2_im;
  temp1_re = y_data[0].re;
  temp1_im = y_data[0].im;
  y_im = y_data[0].re * reconVar1_data[0].im + y_data[0].im * reconVar1_data[0].
    re;
  y_re = y_data[0].re;
  b_y_im = -y_data[0].im;
  y_data[0].re = 0.5 * ((y_data[0].re * reconVar1_data[0].re - y_data[0].im *
    reconVar1_data[0].im) + (y_re * reconVar2_data[0].re - b_y_im *
    reconVar2_data[0].im));
  y_data[0].im = 0.5 * (y_im + (y_re * reconVar2_data[0].im + b_y_im *
    reconVar2_data[0].re));
  y_data[512].re = 0.5 * ((temp1_re * reconVar2_data[0].re - temp1_im *
    reconVar2_data[0].im) + (temp1_re * reconVar1_data[0].re - -temp1_im *
    reconVar1_data[0].im));
  y_data[512].im = 0.5 * ((temp1_re * reconVar2_data[0].im + temp1_im *
    reconVar2_data[0].re) + (temp1_re * reconVar1_data[0].im + -temp1_im *
    reconVar1_data[0].re));
  for (i = 0; i < 255; i++) {
    temp1_re = y_data[i + 1].re;
    temp1_im = y_data[i + 1].im;
    temp2_re_tmp_tmp = wrapIndex_data[i + 1];
    temp2_re_tmp = temp2_re_tmp_tmp - 1;
    temp2_re = y_data[temp2_re_tmp].re;
    temp2_im = y_data[temp2_re_tmp].im;
    y_im = y_data[i + 1].re * reconVar1_data[i + 1].im + y_data[i + 1].im *
      reconVar1_data[i + 1].re;
    y_re = y_data[temp2_re_tmp].re;
    b_y_im = -y_data[temp2_re_tmp].im;
    y_data[i + 1].re = 0.5 * ((y_data[i + 1].re * reconVar1_data[i + 1].re -
      y_data[i + 1].im * reconVar1_data[i + 1].im) + (y_re * reconVar2_data[i +
      1].re - b_y_im * reconVar2_data[i + 1].im));
    y_data[i + 1].im = 0.5 * (y_im + (y_re * reconVar2_data[i + 1].im + b_y_im *
      reconVar2_data[i + 1].re));
    y_data[i + 513].re = 0.5 * ((temp1_re * reconVar2_data[i + 1].re - temp1_im *
      reconVar2_data[i + 1].im) + (temp2_re * reconVar1_data[i + 1].re -
      -temp2_im * reconVar1_data[i + 1].im));
    y_data[i + 513].im = 0.5 * ((temp1_re * reconVar2_data[i + 1].im + temp1_im *
      reconVar2_data[i + 1].re) + (temp2_re * reconVar1_data[i + 1].im +
      -temp2_im * reconVar1_data[i + 1].re));
    y_data[temp2_re_tmp].re = 0.5 * ((temp2_re * reconVar1_data[temp2_re_tmp].re
      - temp2_im * reconVar1_data[temp2_re_tmp].im) + (temp1_re *
      reconVar2_data[temp2_re_tmp].re - -temp1_im * reconVar2_data[temp2_re_tmp]
      .im));
    y_data[temp2_re_tmp].im = 0.5 * ((temp2_re * reconVar1_data[temp2_re_tmp].im
      + temp2_im * reconVar1_data[temp2_re_tmp].re) + (temp1_re *
      reconVar2_data[temp2_re_tmp].im + -temp1_im * reconVar2_data[temp2_re_tmp]
      .re));
    temp2_re_tmp_tmp += 511;
    y_data[temp2_re_tmp_tmp].re = 0.5 * ((temp2_re * reconVar2_data[temp2_re_tmp]
      .re - temp2_im * reconVar2_data[temp2_re_tmp].im) + (temp1_re *
      reconVar1_data[temp2_re_tmp].re - -temp1_im * reconVar1_data[temp2_re_tmp]
      .im));
    y_data[temp2_re_tmp_tmp].im = 0.5 * ((temp2_re * reconVar2_data[temp2_re_tmp]
      .im + temp2_im * reconVar2_data[temp2_re_tmp].re) + (temp1_re *
      reconVar1_data[temp2_re_tmp].im + -temp1_im * reconVar1_data[temp2_re_tmp]
      .re));
  }

  temp1_re = y_data[256].re;
  temp1_im = y_data[256].im;
  y_im = y_data[256].re * reconVar1_data[256].im + y_data[256].im *
    reconVar1_data[256].re;
  y_re = y_data[256].re;
  b_y_im = -y_data[256].im;
  y_data[256].re = 0.5 * ((y_data[256].re * reconVar1_data[256].re - y_data[256]
    .im * reconVar1_data[256].im) + (y_re * reconVar2_data[256].re - b_y_im *
    reconVar2_data[256].im));
  y_data[256].im = 0.5 * (y_im + (y_re * reconVar2_data[256].im + b_y_im *
    reconVar2_data[256].re));
  y_data[768].re = 0.5 * ((temp1_re * reconVar2_data[256].re - temp1_im *
    reconVar2_data[256].im) + (temp1_re * reconVar1_data[256].re - -temp1_im *
    reconVar1_data[256].im));
  y_data[768].im = 0.5 * ((temp1_re * reconVar2_data[256].im + temp1_im *
    reconVar2_data[256].re) + (temp1_re * reconVar1_data[256].im + -temp1_im *
    reconVar1_data[256].re));
}

void c_FFTImplementationCallback_doH(const double x_data[], creal_T y_data[],
  const double costab_data[], const double sintab_data[])
{
  int i;
  int ju;
  int iy;
  double hcostab_data[257];
  double hsintab_data[257];
  int k;
  creal_T reconVar1_data[512];
  int bitrevIndex_data[512];
  creal_T reconVar2_data[512];
  boolean_T tst;
  int wrapIndex_data[512];
  double temp_re;
  double temp_im;
  int iheight;
  int istart;
  int temp_re_tmp;
  int j;
  double twid_re;
  double twid_im;
  int ihi;
  for (i = 0; i < 256; i++) {
    iy = ((i + 1) << 1) - 2;
    hcostab_data[i] = costab_data[iy];
    hsintab_data[i] = sintab_data[iy];
  }

  ju = 0;
  iy = 1;
  for (i = 0; i < 512; i++) {
    reconVar1_data[i].re = sintab_data[i] + 1.0;
    reconVar1_data[i].im = -costab_data[i];
    reconVar2_data[i].re = 1.0 - sintab_data[i];
    reconVar2_data[i].im = costab_data[i];
    if (i + 1 != 1) {
      wrapIndex_data[i] = 513 - i;
    } else {
      wrapIndex_data[0] = 1;
    }

    bitrevIndex_data[i] = 0;
  }

  for (k = 0; k < 511; k++) {
    bitrevIndex_data[k] = iy;
    iy = 512;
    tst = true;
    while (tst) {
      iy >>= 1;
      ju ^= iy;
      tst = ((ju & iy) == 0);
    }

    iy = ju + 1;
  }

  bitrevIndex_data[511] = iy;
  iy = 0;
  for (i = 0; i < 512; i++) {
    y_data[bitrevIndex_data[i] - 1].re = x_data[iy];
    y_data[bitrevIndex_data[i] - 1].im = x_data[iy + 1];
    iy += 2;
  }

  for (i = 0; i <= 510; i += 2) {
    temp_re = y_data[i + 1].re;
    temp_im = y_data[i + 1].im;
    y_data[i + 1].re = y_data[i].re - y_data[i + 1].re;
    y_data[i + 1].im = y_data[i].im - y_data[i + 1].im;
    y_data[i].re += temp_re;
    y_data[i].im += temp_im;
  }

  iy = 2;
  ju = 4;
  k = 128;
  iheight = 509;
  while (k > 0) {
    for (i = 0; i < iheight; i += ju) {
      temp_re_tmp = i + iy;
      temp_re = y_data[temp_re_tmp].re;
      temp_im = y_data[temp_re_tmp].im;
      y_data[temp_re_tmp].re = y_data[i].re - temp_re;
      y_data[temp_re_tmp].im = y_data[i].im - temp_im;
      y_data[i].re += temp_re;
      y_data[i].im += temp_im;
    }

    istart = 1;
    for (j = k; j < 256; j += k) {
      twid_re = hcostab_data[j];
      twid_im = hsintab_data[j];
      i = istart;
      ihi = istart + iheight;
      while (i < ihi) {
        temp_re_tmp = i + iy;
        temp_re = twid_re * y_data[temp_re_tmp].re - twid_im *
          y_data[temp_re_tmp].im;
        temp_im = twid_re * y_data[temp_re_tmp].im + twid_im *
          y_data[temp_re_tmp].re;
        y_data[temp_re_tmp].re = y_data[i].re - temp_re;
        y_data[temp_re_tmp].im = y_data[i].im - temp_im;
        y_data[i].re += temp_re;
        y_data[i].im += temp_im;
        i += ju;
      }

      istart++;
    }

    k /= 2;
    iy = ju;
    ju += ju;
    iheight -= iy;
  }

  c_FFTImplementationCallback_get(y_data, reconVar1_data, reconVar2_data,
    wrapIndex_data);
}

void c_FFTImplementationCallback_gen(double costab_data[], int costab_size[2],
  double sintab_data[], int sintab_size[2], int sintabinv_size[2])
{
  double costab1q_data[257];
  int k;
  double costab_tmp;
  double sintab_tmp;
  costab1q_data[0] = 1.0;
  for (k = 0; k < 128; k++) {
    costab1q_data[k + 1] = cos(0.0061359231515425647 * ((double)k + 1.0));
  }

  for (k = 0; k < 127; k++) {
    costab1q_data[k + 129] = sin(0.0061359231515425647 * (256.0 - ((double)k +
      129.0)));
  }

  costab1q_data[256] = 0.0;
  costab_size[0] = 1;
  costab_size[1] = 513;
  sintab_size[0] = 1;
  sintab_size[1] = 513;
  costab_data[0] = 1.0;
  sintab_data[0] = 0.0;
  for (k = 0; k < 256; k++) {
    costab_tmp = costab1q_data[k + 1];
    costab_data[k + 1] = costab_tmp;
    sintab_tmp = -costab1q_data[255 - k];
    sintab_data[k + 1] = sintab_tmp;
    costab_data[k + 257] = sintab_tmp;
    sintab_data[k + 257] = -costab_tmp;
  }

  sintabinv_size[0] = 1;
  sintabinv_size[1] = 0;
}

/* End of code generation (FFTImplementationCallback.c) */
