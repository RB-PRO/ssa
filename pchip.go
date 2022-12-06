package main

import (
	"math"
)

func exteriorSlope(d1, d2, h1, h2 float64) float64 {
	var s float64
	var signd1 float64
	var signs float64
	s = ((2.0*h1+h2)*d1 - h1*d2) / (h1 + h2)
	signd1 = d1
	if d1 < 0.0 {
		signd1 = -1.0
	} else if d1 > 0.0 {
		signd1 = 1.0
	} else {
		if d1 == 0.0 {
			signd1 = 0.0
		}
	}
	signs = s
	if s < 0.0 {
		signs = -1.0
	} else if s > 0.0 {
		signs = 1.0
	} else {
		if s == 0.0 {
			signs = 0.0
		}
	}
	if signs != signd1 {
		s = 0.0
	} else {
		signs = d2
		if d2 < 0.0 {
			signs = -1.0
		} else if d2 > 0.0 {
			signs = 1.0
		} else {
			if d2 == 0.0 {
				signs = 0.0
			}
		}
		if (signd1 != signs) && (math.Abs(s) > math.Abs(3.0*d1)) {
			s = 3.0 * d1
		}
	}
	return s
}
func pchip(x, y, new_x []float64, x_len, new_x_len int) []float64 {
	new_y := make([]float64, new_x_len)
	var low_ip1 int
	var hs float64
	del := make([]float64, x_len-1)
	slopes := make([]float64, x_len)
	h := make([]float64, x_len-1)
	var hs3 float64
	var w1 float64
	//var ix int
	pp_coefs := make([]float64, (x_len-1)+(3*(x_len-1)))
	var low_i int
	var high_i int
	var mid_i int
	for low_ip1 := 0; low_ip1 < x_len-1; low_ip1++ {
		hs = x[low_ip1+1] - x[low_ip1]
		del[low_ip1] = (y[low_ip1+1] - y[low_ip1]) / hs
		h[low_ip1] = hs
	}

	for low_ip1 := 0; low_ip1 < x_len-2; low_ip1++ {
		hs = h[low_ip1] + h[low_ip1+1]
		hs3 = 3.0 * hs
		w1 = (h[low_ip1] + hs) / hs3
		hs = (h[low_ip1+1] + hs) / hs3
		hs3 = 0.0
		if del[low_ip1] < 0.0 {
			if del[low_ip1+1] <= del[low_ip1] {
				hs3 = del[low_ip1] / (w1*(del[low_ip1]/del[low_ip1+1]) + hs)
			} else {
				if del[low_ip1+1] < 0.0 {
					hs3 = del[low_ip1+1] / (w1 + hs*(del[low_ip1+1]/del[low_ip1]))
				}
			}
		} else {
			if del[low_ip1] > 0.0 {
				if del[low_ip1+1] >= del[low_ip1] {
					hs3 = del[low_ip1] / (w1*(del[low_ip1]/del[low_ip1+1]) + hs)
				} else {
					if del[low_ip1+1] > 0.0 {
						hs3 = del[low_ip1+1] / (w1 + hs*(del[low_ip1+1]/del[low_ip1]))
					}
				}
			}
		}

		slopes[low_ip1+1] = hs3
	}

	slopes[0] = exteriorSlope(del[0], del[1], h[0], h[1])
	slopes[x_len-1] = exteriorSlope(del[x_len-2], del[x_len-3], h[x_len-2], h[x_len-3])
	for low_ip1 := 0; low_ip1 < x_len-1; low_ip1++ {
		hs = (del[low_ip1] - slopes[low_ip1]) / h[low_ip1]
		hs3 = (slopes[low_ip1+1] - del[low_ip1]) / h[low_ip1]
		pp_coefs[low_ip1] = (hs3 - hs) / h[low_ip1]
		pp_coefs[low_ip1+x_len-1] = 2.0*hs - hs3
		pp_coefs[low_ip1+(2*(x_len-1))] = slopes[low_ip1]
		pp_coefs[low_ip1+(3*(x_len-1))] = y[low_ip1]
	}

	for ix := 0; ix < new_x_len; ix++ {
		low_i = 0
		low_ip1 = 2
		high_i = x_len
		for high_i > low_ip1 {
			mid_i = ((low_i + high_i) + 1) >> 1
			if new_x[ix] >= x[mid_i-1] {
				low_i = mid_i - 1
				low_ip1 = mid_i + 1
			} else {
				high_i = mid_i
			}
		}

		hs = new_x[ix] - x[low_i]
		hs3 = pp_coefs[low_i]
		for low_ip1 := 0; low_ip1 < 3; low_ip1++ {
			hs3 = hs*hs3 + pp_coefs[low_i+(low_ip1+1)*(x_len-1)]
		}

		new_y[ix] = hs3
		/*
			// my
			if new_y[ix] < 0.15 {
				new_y[ix] = new_y[ix-1]
				fmt.Println("YESS")
			}
		*/
	}
	return new_y
}

// Функция, реализующая сглаживание с помощью фильтра Савицкого-Голея

// Вторая пробная реализация
func pchip2(x, y, t []float64, x_len, new_x_len int) []float64 {
	sizex,sizexy=x_len
	
	
v:=make([]float64,x_len)
h:=make([]float64,x_len)
del:=make([]float64,x_len)

  for k:=0;k<x_len;k++ {
    h[k] = x[k+1]-x.at[k]
    del[k] = (y[k+1]-y[k])/h[k]

  }



    var m float64= sizey+1;
    //std::vector<double> slopes,dzzdx,dzdxdx;
    slopes:=pchipslopes(x, y, del)
    double n = sizex+1;
    double d = 1;
    std::vector<double> dxd = h;
    dzzdx.resize(slopes.size()-1);
    dzdxdx.resize(slopes.size()-1);

    for(int k=0;k<slopes.size()-1;k++){


      dzzdx.at(k) = (del.at(k)-slopes.at(k))/dxd.at(k);
      dzdxdx.at(k) = (slopes.at(k+1)-del.at(k))/dxd.at(k);


    }
    

    int dm1 = (n-1)*d;
    Eigen::MatrixXd coeff(dm1,4);
    
    for(int j=0;j<dm1;j++)
    {
      coeff.row(j) << (dzdxdx.at(j)-dzzdx.at(j))/dxd.at(j), 2*dzzdx.at(j)-dzdxdx.at(j), slopes.at(j),y.at(j);
    }

    int l = 1;
    v.resize(t.size());
    for(int j=0;j<t.size()-1;j++){
      if(t.at(j) < x.at(l)) {
        double t_aux = t.at(j)-x.at(l-1);
        v.at(j) = coeff(l-1,0)*pow(t_aux,3)+coeff(l-1,1)*pow(t_aux,2)+coeff(l-1,2)*pow(t_aux,1)+coeff(l-1,3);

      }
      else{
        l++;
        v.at(j) = coeff(l-1,3);

      } 
  }
        v.at(v.size()-1) = y.at(y.size()-1);
	return []float64{}
}
func pchipslopes(x, y, del []float64)[]float64{


	var n = len(x);
    var size = len(y);

	d:=make([]float64,size)


  if(n==2) {
    for i:=0;i<size;i++ {
		d[i] = del[0]}
    return d
  }

  std::vector<int> index;

   var found bool= false

  for  k:=0;k<size-2;k++  {
    var check float64 = sign(del.at(k))*sign(del.at(k+1))
    if(check > 0){

      index.push_back(k);
      found = true;
    }
  }

  std::vector<double> h;
  h.resize(size-1);
  
  


   // hs = h(k)+h(k+1);
   // w1 = (h(k)+hs)./(3*hs);
   // w2 = (hs+h(k+1))./(3*hs);
   // dmax = max(abs(del(k)), abs(del(k+1)));
   // dmin = min(abs(del(k)), abs(del(k+1)));
   // d(k+1) = dmin./conj(w1.*(del(k)./dmax) + w2.*(del(k+1)./dmax));

  for(int k=0;k<size-1;k++)
  {
    h.at(k) = x.at(k+1)-x.at(k);

  }
    if(found){
    for(int j=0;j<index.size();j++){

    double hs = h.at(index.at(j))+h.at(index.at(j)+1);
    double w1 = (h.at(index.at(j))+hs)/(3*hs);
    double w2 = (h.at(index.at(j)+1)+hs)/(3*hs);
    double dmax = std::max(fabs(del.at(index.at(j))),fabs(del.at(index.at(j)+1)));
    double dmin = std::min(fabs(del.at(index.at(j))),fabs(del.at(index.at(j)+1)));
    d.at(index.at(j)+1) = dmin/(w1*del.at(index.at(j))/dmax + w1*del.at(index.at(j)+1)/dmax);
    }
  }
    
 
   d.at(0) = ((2*h.at(0)+h.at(1))*del.at(0) - h.at(0)*del.at(1))/(h.at(0)+h.at(1));
   if (sign(d.at(0)) != sign(del.at(0)))
   {
      d.at(0) = 0;
   }
   else if ((sign(del.at(0)) != sign(del.at(1))) && (fabs(d.at(0)) > fabs(3*del.at(0))))
   {
      d.at(0) = 3*del.at(0);
   }
   
   d.at(n-1) = ((2*h.at(n-2)+h.at(n-3))*del.at(n-2) - h.at(n-2)*del.at(n-3))/(h.at(n-2)+h.at(n-3));
   if(sign(d.at(n-1)) != sign(del.at(n-2)))
   { 
      d.at(n-1) = 0;
   }
   else if ((sign(del.at(n-2)) != sign(del.at(n-3))) && (fabs(d.at(n-1)) > fabs(3*del.at(n-2))))
   {
      d.at(n-1) = 3*del.at(n-2);
   }
}