func (c *consumer) setLeader(isLeader bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isLeader == isLeader {
		return
	}

	c.isLeader = isLeader

	if isLeader {
		// Ensure ACK subscription is active when becoming leader
		if c.ackSub == nil {
			c.subscribeToAcks()
		}
	} else {
		// Cleanup ACK subscription when stepping down
		if c.ackSub != nil {
			c.ackSub.Unsubscribe()
			c.ackSub = nil
		}
	}
}