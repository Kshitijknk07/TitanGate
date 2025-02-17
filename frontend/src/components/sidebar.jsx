"use client"

import * as React from "react"
import { motion } from "framer-motion"
import { ChevronRight, LayoutDashboard, Activity, Network, Database } from 'lucide-react'

import { cn } from "../libs/utils"
import { Button } from "@/components/ui/button"
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

const menuItems = [
  { name: "API Analytics", icon: Activity },
  { name: "Load Balancing", icon: Network },
  { name: "GraphQL Gateway", icon: Database },
]

export function Sidebar() {
  const [isOpen, setIsOpen] = React.useState(true)

  return (
    <TooltipProvider>
      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
        className="fixed left-0 top-0 z-40 h-screen"
      >
        <motion.div
          initial={{ width: 64 }}
          animate={{ width: isOpen ? 250 : 64 }}
          transition={{ duration: 0.3, ease: "easeInOut" }}
          className="flex h-full flex-col justify-between border-r border-gray-200 bg-white shadow-sm"
        >
          <div className="flex flex-col space-y-6 p-4">
            <CollapsibleTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="self-end"
                aria-label="Toggle sidebar"
              >
                <ChevronRight
                  className={cn(
                    "h-4 w-4 transition-transform duration-200",
                    isOpen && "rotate-180"
                  )}
                />
              </Button>
            </CollapsibleTrigger>
            <motion.div
              animate={{ opacity: isOpen ? 1 : 0 }}
              transition={{ duration: 0.2 }}
            >
              <h2 className="px-2 text-lg font-semibold tracking-tight">
                Navigation
              </h2>
            </motion.div>
            <nav className="space-y-2">
              {menuItems.map((item, index) => (
                <Tooltip key={item.name}>
                  <TooltipTrigger asChild>
                    <motion.div
                      whileHover={{ x: 5 }}
                      transition={{ duration: 0.2 }}
                    >
                      <Button
                        variant="ghost"
                        className={cn(
                          "w-full justify-start",
                          !isOpen && "px-2"
                        )}
                      >
                        <item.icon className="mr-2 h-4 w-4" />
                        <span
                          className={cn(
                            "transition-opacity",
                            isOpen ? "opacity-100" : "opacity-0"
                          )}
                        >
                          {item.name}
                        </span>
                      </Button>
                    </motion.div>
                  </TooltipTrigger>
                  <TooltipContent side="right" sideOffset={10}>
                    {item.name}
                  </TooltipContent>
                </Tooltip>
              ))}
            </nav>
          </div>
          <div className="p-4">
            <Button variant="outline" className="w-full">
              <LayoutDashboard className="mr-2 h-4 w-4" />
              <span
                className={cn(
                  "transition-opacity",
                  isOpen ? "opacity-100" : "opacity-0"
                )}
              >
                Dashboard
              </span>
            </Button>
          </div>
        </motion.div>
      </Collapsible>
    </TooltipProvider>
  )
}