#!/usr/bin/env python
###############################################################################
#  Copyright (C) 1994 - 2009, Performance Dynamics Company                    #
#                                                                             #
#  This software is licensed as described in the file COPYING, which          #
#  you should have received as part of this distribution. The terms           #
#  are also available at http://www.perfdynamics.com/Tools/copyright.html.    #
#                                                                             #
#  You may opt to use, copy, modify, merge, publish, distribute and/or sell   #
#  copies of the Software, and permit persons to whom the Software is         #
#  furnished to do so, under the terms of the COPYING file.                   #
#                                                                             #
#  This software is distributed on an "AS IS" basis, WITHOUT WARRANTY OF ANY  #
#  KIND, either express or implied.                                           #
###############################################################################

#
#  $Id: test.py,v 4.4 2009/03/26 02:55:33 pfeller Exp $
#
#---------------------------------------------------------------------

import pdq

#---- Define globals -------------------------------------------------

arrivRate    = 0.75
service_time = 1.0

#---- Initialize -----------------------------------------------------

pdq.Init("OpenCenter")
pdq.SetComment("A simple M/M/1 queue")



#---- Define the workload and circuit type ---------------------------

pdq.streams = pdq.CreateOpen("work", arrivRate)

pdq.SetWUnit("Customers")
pdq.SetTUnit("Seconds")


#---- Define the queueing center -------------------------------------

pdq.nodes = pdq.CreateNode("server", pdq.CEN, pdq.FCFS)

#---- Define service demand due to workload on the queueing center ---

pdq.SetDemand("server", "work", service_time)

#---- Solve the model ------------------------------------------------
#  Must use the CANONical method for an open circuit

pdq.Solve(pdq.CANON)

#---- Generate a report ----------------------------------------------

pdq.Report()

#---------------------------------------------------------------------

comment = pdq.GetComment()

print 'pdq.GetComment -> \"%s\"' % comment

#---------------------------------------------------------------------

