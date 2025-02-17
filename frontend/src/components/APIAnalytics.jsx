"use client"

import * as React from "react"
import { motion } from "framer-motion"
import { Activity, AlertCircle, BarChart } from 'lucide-react'

import { cn } from "@/lib/utils"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"

const API_URL = process.env.NEXT_PUBLIC_API_URL

export function APIAnalytics() {
  const [traffic, setTraffic] = React.useState(0)
  const [errors, setErrors] = React.useState(0)

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`${API_URL}/metrics`)
        if (!response.ok) throw new Error('Failed to fetch metrics')
        const data = await response.json()
        setTraffic(data.traffic)
        setErrors(data.errors)
      } catch (error) {
        console.error("Error fetching data:", error)
        setTraffic(0)
        setErrors(0)
      }
    }

    fetchData()
    
    const interval = setInterval(fetchData, 30000)
    return () => clearInterval(interval)
  }, [])

  const errorRate = errors / traffic * 100
  const errorRateFormatted = errorRate.toFixed(2)

  return (
    <motion.div 
      initial={{ opacity: 0, y: 20 }} 
      animate={{ opacity: 1, y: 0 }} 
      transition={{ duration: 0.5 }}
      className="w-full max-w-2xl mx-auto"
    >
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl font-bold">API Analytics</CardTitle>
          <CardDescription>Track request metrics, performance, and errors in real-time.</CardDescription>
        </CardHeader>
        <CardContent className="grid gap-6">
          <motion.div 
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2, duration: 0.5 }}
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Activity className="h-4 w-4 text-blue-500" />
                <h4 className="font-semibold">Requests</h4>
              </div>
              <span className="text-2xl font-bold">{traffic.toLocaleString()}</span>
            </div>
            <Progress value={100} className="mt-2" />
          </motion.div>
          <motion.div 
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.4, duration: 0.5 }}
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <AlertCircle className="h-4 w-4 text-red-500" />
                <h4 className="font-semibold">Errors</h4>
              </div>
              <span className="text-2xl font-bold">{errors.toLocaleString()}</span>
            </div>
            <Progress value={(errors / traffic) * 100} className="mt-2" />
          </motion.div>
          <motion.div 
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.6, duration: 0.5 }}
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <BarChart className="h-4 w-4 text-green-500" />
                <h4 className="font-semibold">Error Rate</h4>
              </div>
              <span className={cn(
                "text-2xl font-bold",
                errorRate > 1 ? "text-red-500" : "text-green-500"
              )}>
                {errorRateFormatted}%
              </span>
            </div>
            <Progress 
              value={errorRate} 
              className={cn(
                "mt-2",
                errorRate > 1 ? "text-red-500" : "text-green-500"
              )} 
            />
          </motion.div>
        </CardContent>
      </Card>
    </motion.div>
  )
}
