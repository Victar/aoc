package adventofcode.year2021;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day17 extends BaseTest {

	static int xMin = 81;
	static int xMax = 129;
	static int yMin = -150;
	static int yMax = -108;

	//	int xMin = 20;
	//	int xMax = 30;
	//	int yMin = -10;
	//	int yMax = -5;

	@Test
	public void runBoth() throws Exception {
		int maxY = 0;
		int count = 0;
		for (int x = -1000; x < 1000; x++) {
			for (int y = -1000; y < 1000; y++) {
				final Velocity current = new Velocity(x, y);
				int currentX = 0;
				int currentY = 0;
				int currentYMax = 0;

				boolean hit = false;
				boolean check = true;
				while (check) {
					currentYMax = Math.max(currentYMax, currentY);
					hit = hit || isHit(currentX, currentY);
					check = inArea(currentX, currentY);
					currentX = current.x + currentX;
					currentY = current.y + currentY;
					current.updateVelocity();

				}
				if (hit) {
					count++;
					maxY = Math.max(currentYMax, maxY);
				}
			}
		}
		System.out.println(maxY);
		System.out.println(count);

	}

	public boolean isHit(final int currentX, final int currentY) {
		return (this.xMin <= currentX) && (currentX <= this.xMax) && (this.yMin <= currentY) && (currentY <= this.yMax);
	}

	public boolean inArea(final int currentX, final int currentY) {
		return currentY >= this.yMin;
	}

	@Data public class Velocity {

		int x;
		int y;

		public Velocity(final int x, final int y) {
			this.x = x;
			this.y = y;
		}

		//	Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
		//	Due to gravity, the probe's y velocity decreases by 1
		public void updateVelocity() {
			this.y = this.y - 1;
			if (this.x > 0) {
				this.x = Math.max(this.x - 1, 0);
			} else {
				this.x = Math.min(this.x + 1, 0);
			}
		}
	}

}
