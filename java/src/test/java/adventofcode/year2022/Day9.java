package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day9 extends BaseTest {

	public static final int DAY = 9;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		runAny(1);
	}

	@Test public void runGold() throws Exception {
		runAny(9);
	}

	public void runAny(int size) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		Rope rope = new Rope(size + 1);
		for (final String input : data) {
			rope.doInstructionAny(input, size);
		}
		System.out.println(rope.getVisited().size());
	}

	@Data static class Rope {

		Set<String> visited = new HashSet<>();
		Point head = new Point();
		List<Point> tail = new ArrayList<>();

		Point moveU = new Point(0, -1);
		Point moveD = new Point(0, 1);
		Point moveL = new Point(-1, 0);
		Point moveR = new Point(1, 0);

		Map<String, Point> moves = new HashMap<>();

		Rope(int length) {
			for (int i = 0; i < length; i++) {
				tail.add(new Point());
			}
			this.moves.put("U", moveU);
			this.moves.put("D", moveD);
			this.moves.put("L", moveL);
			this.moves.put("R", moveR);
		}

		public void doInstructionAny(String input, int count) {
			String[] arr = input.split(" ");
			int length = Integer.parseInt(arr[1]);
			Point move = moves.get(arr[0]);

			for (int i = 0; i < length; i++) {
				tail.set(0, tail.get(0).add(move));
				for (int j = 1; j < tail.size(); j++) {
					Point newTail = getNewTail(tail.get(j - 1), tail.get(j));
					tail.set(j, newTail);
				}
				visited.add(tail.get(count).getId());
			}
		}

		public Point getNewTail(Point head, Point t) {
			Point tail = new Point(t.x, t.y);
			int distX = head.x - tail.x;
			int distY = head.y - tail.y;
			if (Math.abs(distX) > 1 || Math.abs(distY) > 1) {
				if (Math.abs(distX) != 0) {
					tail.x = tail.x + (distX > 0 ? 1 : -1);
				}
				if (Math.abs(distY) != 0) {
					tail.y = tail.y + (distY > 0 ? 1 : -1);
				}
			}
			return tail;
		}
	}

	@Data static class Point {

		int x = 0;
		int y = 0;

		public Point() {
		}

		public Point(int x, int y) {
			this.x = x;
			this.y = y;
		}

		public String getId() {
			return x + "-" + y;
		}

		public Point add(Point move) {
			return new Point(this.x + move.x, this.y + move.y);
		}
	}

}
